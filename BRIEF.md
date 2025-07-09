# Registry Take-Home

## Task

We have provided you with a repository containing an API implemented in golang. This API contains endpoints returning modules and templates that can be used within Coder to define new Cloud Development Environments for the users. For contextual information (not relevant to the challenge implementation, but to help you understand the background here):

- [**A template**](https://coder.com/docs/admin/templates) is the root Terraform configuration that Coder executes every time it creates, starts, or stops a workspace. It declares the infrastructure (VM vs Kubernetes pod vs Docker), the base image, resource limits, parameters.
- [**A module**](https://coder.com/docs/admin/templates/extending-templates/modules) is a re-usable Terraform module that you import inside a template's main.tf. Think of it as a Lego brick that layers extra capability on top of the workspace the template already provisions.

Both templates and modules are characterized by the following fields:

- name
- description
- logo (URI to the image)
- contributor - think of it as a github username (e.g. `coder`, `johndoe`, etc.)
- operating system - union type - `'Windows' | 'Linux' | 'MacOS'`
- source - union type - `'Partner' | 'Official'`
- custom tags - custom strings that have any alphanumeric value (e.g. `ai`, `rdp`, `gcp`, `aws`)

Your task is to create a UI in TypeScript+React that has two client-rendered pages, one for templates and one for modules. For each page, the UI should:

- Display a list of resources associated with the page (i.e., templates for a `/templates` route, modules for a `/modules` route).
- Use a provided backend API to search through each resource type via auto-complete.
- Give the user a way to clear the selected predicates, making sure that the UI for clearing the search elements is uniquely associated with the other UI controls.
- Use a provided backend API to display notifications to the user as a new resource gets published for the current page. Once a user navigates away from a page for a given resource type, they should no longer receive notifications for it.
- Add UI to allow the user to delete modules (that feature will not be available on the templates page).

You may use any additional state management libraries, but it's expected that you can solve the challenge with just React and React Router.

## **Submission**

1. Share your submission with us as a Pull Request to your forked GitHub repo
2. Add the following Coder team members as collaborators:
   - @spikecurtis
   - @ibetitsmike
   - @Parkreiner
3. When filling out the PR message, feel free to leave any additional notes or comments that you feel you would us review the code. This can be things like why one specific choice was made, or things that you wish you could've done if you had more time. We are not looking for a perfect submission – we are looking for signs of how you approach engineering thinking.

## Backend API documentation

### Features

- List and filter modules and templates
- Delete modules
- Auto-complete module and template names
- Real-time updates via Server-Sent Events (SSE)
- Background daemon that periodically adds random resources

### API Endpoints

- `GET /modules` - List all modules (optional query param: `name` for filtering)
- `GET /templates` - List all templates (optional query param: `name` for filtering)
- `GET /autocomplete/modules` - Get module name suggestions (query param: `prefix`)
- `GET /autocomplete/templates` - Get template name suggestions (query param: `prefix`)
- `DELETE /modules/{id}` - Delete a module by ID
- `DELETE /templates/{id}` - Delete a template by ID
- `GET /events` - SSE endpoint for real-time updates

### Usage

#### Running the Server Locally

```bash
go run ./main.go
```

The server will start on port 8080 by default.

### Testing

Run tests with:

```bash
go test ./...
```

#### Example Requests

List all modules:

```bash
curl http://localhost:8080/modules
```

Filter modules by name:

```bash
curl http://localhost:8080/modules?name=awesome
```

Delete a module by ID:

```bash
curl -X DELETE http://localhost:8080/modules/{module-id}
```

Delete a template by ID:

```bash
curl -X DELETE http://localhost:8080/templates/{template-id}
```

Listen for real-time updates:

```bash
curl -N http://localhost:8080/events
```

### Implementation Details

The project consists of a `main.go` file and a `server` package that contains:

- A background daemon that adds new modules and templates on an interval.
- An in-memory database implementation.
- An HTTP server with request handlers for modules and templates.

## Help with getting started on the frontend

This take-home deliberately gives you a lot of freedom as far as how you want to bootstrap a new project. You are welcome to use any libraries or tools (including starter templates created by someone else), as long as you stay within the TypeScript+React ecosystem. No matter what you choose, though, you will be expected to explain every line of code you submit.

That said, here are some tools that we use at Coder that you may find useful.

### Libraries

- [Vite](https://vite.dev/)
- [React Router](https://reactrouter.com/) (SPA-style rendering is perfectly fine; we don't need you to do any server-rendering unless you really want to)
- The [TanStack ecosystem](https://tanstack.com/), particularly [TanStack Query](https://tanstack.com/query/latest)
- [Radix UI Primitives](https://www.radix-ui.com/primitives/docs/overview/introduction)
- [TailwindCSS](https://tailwindcss.com/)

### Utility functions

#### `useEffectEvent`

This is a polyfill for the `useEffectEvent` hook that the React team proposed around 2023. They have still not added it to the core library, but we have found it incredibly useful for wrangling `useEffect` code in our production apps. You might find it helpful as well.

Note that if you use this custom hook, your take-home reviewer may ask you about the hook, as far as:

1. What problem the hook solves
2. How the hook interacts with React's life cycle, and why it's implemented with these specific built-in hooks
3. How the TypeScript types in the hook work

```ts
import { useCallback, useLayoutEffect, useRef } from "react";

/**
 * @see {@link https://react.dev/reference/react/experimental_useEffectEvent}
 */
export function useEffectEvent<TArgs extends unknown[], TReturn = unknown>(
  callback: (...args: TArgs) => TReturn
) {
  const callbackRef = useRef(callback);
  useLayoutEffect(() => {
    callbackRef.current = callback;
  }, [callback]);

  return useCallback((...args: TArgs): TReturn => {
    return callbackRef.current(...args);
  }, []);
}
```
