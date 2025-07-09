
# ARCHITECTURE.md

## Overview

This document outlines the architecture, workflow, and process for the Coder Registry UI take-home assignment. The approach is incremental and feature-based, iterating towards an MVP with a focus on clarity, maintainability, and real-world engineering practices.

---

## Tech Stack & Tooling

**Core Stack:**
- **Vite** + **React** + **TypeScript**: Fast, modern SPA development
- **React Router**: Client-side routing
- **Tailwind CSS** + **Flowbite**: Utility-first styling and design system

**State Management & Data:**
- **TanStack Query**: Data fetching and caching
- **Zustand**: Lightweight, global state management

**Development & Quality:**
- **ESLint**, **Prettier**, **Husky**: Code quality and pre-commit checks
- **DevContainers**: Consistent development environment
- **Storybook**: Component-driven development (documented, not implemented)
- **Jest**: Unit testing for non-UI logic (documented, not implemented)

---

## Architecture Patterns

**Core Patterns:**
- **Container/Presentational Pattern:** Containers handle data fetching and logic; presentational components handle rendering and UI
- **Feature-Based Structure:** Each feature encapsulated in its own directory with components, containers, logic, tests, and stories
- **Custom Hooks:** Reusable logic encapsulation (e.g., useTemplates, useModules, useNotifications)

**State Management Strategy:**
- **Server State:** Managed by TanStack Query (fetching, caching, mutations, background refetching)
- **UI/Local State:** Managed by Zustand (toasts, modals, theme)
- **Notifications:** Context-based system for cross-component communication

**Performance & UX Patterns:**
- **Error Boundaries:** Graceful error handling and display
- **Loading States:** Spinners and skeletons during data fetching
- **Debounced Search:** Optimized API calls with race condition prevention
- **Imperative Cache Updates:** SSE events trigger real-time UI synchronization

---

## Directory Structure

The project uses a **feature-based directory structure**. Each feature (mapped to an epic or major user-facing capability) is encapsulated in its own directory under `src/`, containing all related source files: presentational components, containers, logic, tests, stories, and an `index.ts` for exports.

```
src/
│
├── navigation/           # Persistent top bar and navigation links
│   ├── components/
│   │   └── NavigationBar.tsx
│   ├── containers/
│   │   └── NavigationBarContainer.tsx
│   ├── logic/
│   ├── stories/
│   ├── tests/
│   └── index.ts
│
├── search/               # Search UI and logic (used in top bar, but feature-encapsulated)
│   ├── components/
│   │   └── SearchInput.tsx
│   ├── containers/
│   │   └── SearchInputContainer.tsx
│   ├── logic/
│   ├── stories/
│   ├── tests/
│   └── index.ts
│
├── templates/            # Templates list and related logic
│   ├── components/
│   │   ├── TemplateList.tsx
│   │   └── TemplateCard.tsx
│   ├── containers/
│   │   └── TemplateListContainer.tsx
│   ├── logic/
│   ├── stories/
│   ├── tests/
│   └── index.ts
│
├── modules/              # Modules list, delete, and related logic
│   ├── components/
│   │   ├── ModuleList.tsx
│   │   └── ModuleCard.tsx
│   ├── containers/
│   │   └── ModuleListContainer.tsx
│   ├── logic/
│   ├── stories/
│   ├── tests/
│   └── index.ts
│
├── notifications/        # Toasts and real-time notifications (context-aware)
│   ├── components/
│   │   └── Notification.tsx
│   ├── containers/
│   │   └── NotificationProvider.tsx
│   ├── logic/
│   ├── stories/
│   ├── tests/
│   └── index.ts
│
├── common/               # Shared UI components, styles, logic, and utilities
│   ├── components/
│   ├── logic/
│   ├── utils/
│   ├── styles/
│   │   ├── tailwind.config.js
│   │   ├── flowbite.config.js
│   │   ├── index.css
│   │   └── ...
│   └── index.ts
│
├── layouts/              # Layout components for persistent UI (e.g., top bar)
│   └── MainLayout.tsx
│
├── views/                # Route-level views (pages), each wraps content with MainLayout
│   ├── TemplatesView.tsx
│   ├── ModulesView.tsx
│   └── ...
│
├── App.tsx               # App shell, route definitions
├── main.tsx              # Entry point
└── types/                # Global TypeScript types/interfaces
    └── index.ts
```

### Feature Encapsulation, Views, and Layouts

- **Layouts:**  
  Layout components (e.g., `MainLayout`) encapsulate persistent UI such as the navigation bar. Views are rendered as children, ensuring shared UI is not unmounted between route changes. This improves performance and code clarity.
- **Views:**  
  Each route-level view (e.g., `TemplatesView`, `ModulesView`) is responsible for wrapping its content with the appropriate layout (typically `MainLayout`). This allows for flexible layout composition and keeps routing logic clean. Views compose feature components (e.g., navigation, search, templates/modules list, notifications) as needed.
- **Navigation and Search:**  
  The navigation bar is a persistent UI element, implemented in `navigation/`. The search feature is encapsulated in its own directory (`search/`) and imported into the navigation bar as a child component. This keeps search logic and tests isolated, while allowing for easy composition in the UI.
- **Notifications:**  
  The notification system is encapsulated in `notifications/` and provided via context at the app level. Feature-specific notifications (e.g., for templates or modules) are triggered by the relevant feature logic but rendered by the shared notification provider.
- **Common Components:**  
  Shared UI (e.g., spinners, error boundaries) and utilities live in `common/` and are imported as needed by features.

### Rationale

- **Scalability:**  
  Each feature is self-contained, making it easy to add, refactor, or remove features without impacting unrelated code.
- **Testability:**  
  Tests and stories are colocated with their feature, ensuring they stay up-to-date and easy to discover.
- **Separation of Concerns:**  
  Views handle page-level composition and layout, while features encapsulate reusable logic/UI. Cross-feature UI (like search in the nav bar) is handled via composition, not by merging feature directories.
- **Performance:**  
  Layouts prevent unnecessary unmounting of persistent UI (e.g., top bar), improving user experience and efficiency.
- **Clarity for Reviewers:**  
  The structure maps directly to user-facing features and epics, making it easy for reviewers to navigate and understand the codebase.

---

## User Journeys & Data Flow

### 1. View Templates
- User navigates to the Templates page.
- The Templates container uses TanStack Query's `useQuery` to fetch templates from the backend.
- While loading, a spinner or skeleton is shown.
- On success, the presentational component renders the list of templates.
- On error, an error boundary or message is displayed.

### 2. View Modules
- User navigates to the Modules page.
- The Modules container uses TanStack Query's `useQuery` to fetch modules from the backend.
- Loading, error, and data states are handled as above.

### 3. Delete a Module
- User clicks delete on a module card.
- The container calls TanStack Query's `useMutation` to delete the module.
- On success, the modules query is invalidated, triggering a refetch and UI update.
- A toast notification is shown for feedback.

### 4. Search Templates or Modules
- User types in the search input (debounced).
- The search container uses TanStack Query's `useQuery` to fetch autocomplete suggestions.
- Suggestions are shown in a dropdown.
- Selecting a suggestion or submitting triggers a filtered fetch of templates/modules.

### 5. Real-Time Notification & UI Sync
- App opens an SSE connection to `/events`.
- When a new template/module is published or deleted, the backend sends an event.
- The frontend receives the event, shows a toast notification, and imperatively invalidates the relevant TanStack Query cache.
- TanStack Query refetches the latest data, and the UI updates in real time.

---

## Feature-Based Development Workflow

To ensure clarity, testability, and incremental delivery, features are organized as vertical epics in GitHub — each representing a specific domain capability.

Each epic contains a linear (waterfall) progression of atomic tasks:

1. Setup any additional dependencies or tooling
2. Write business logic (e.g. Zustand slice, utilities)
3. Compose UI components (with Storybook support)
4. Add unit and interaction tests
5. Add E2E tests (Cypress) if needed
6. Raise PR and release

This structure enables:

- **Iterative release** — Core functionality can ship early, with enhancements layered in later
- **Atomic PRs** — Each task delivers clear value and can be reviewed/tested in isolation
- **Tight alignment with architecture** — Feature folders map 1:1 with epics and their tasks
- **Parallel progress** — UI, logic, and tests can move forward in small units without waiting for an entire feature to be done
- **Low technical debt** — Features are validated end-to-end before moving on

This approach mirrors the structure of the file system and promotes maintainability as the app scales.

---

## API Endpoints (Reference)

- `GET /modules` - List all modules (optional query param: `name` for filtering)
- `GET /templates` - List all templates (optional query param: `name` for filtering)
- `GET /autocomplete/modules` - Get module name suggestions (query param: `prefix`)
- `GET /autocomplete/templates` - Get template name suggestions (query param: `prefix`)
- `DELETE /modules/{id}` - Delete a module by ID
- `GET /events` - SSE endpoint for real-time updates

---

## Trade-offs & Design Decisions

- **TanStack Query over Redux/MobX:** Eliminates boilerplate and simplifies data fetching/caching patterns
- **Feature-based architecture:** API logic colocated with features for modularity; can be promoted to `/common` if shared
- **SPA over SSR/SSG:** No Next.js/Remix as requirements are SPA-focused
- **Zustand for UI state:** Lightweight solution that covers all UI state needs
- **Time-boxed scope:** No mobile support or persistent notification logs due to time constraints

---

## Quality & Testing Strategy

- **Unit Tests:** For logic and utility functions (Jest, documented but not implemented)
- **Component Tests:** For presentational components (Storybook test runner, documented but not implemented)
- **E2E Tests:** For user journeys (documented but not implemented)
- **Visual Regression:** Storybook for UI review and regression (documented but not implemented)
- **Accessibility:** Part of definition of done for UI features (documented, not fully implemented)
- **Responsiveness:** Desktop-first, but layout and components designed to be responsive using Flowbite and Tailwind

---

## Future Enhancements

- **Error Handling:** Comprehensive error handling and user-friendly error UI
- **Test Coverage:** Full test coverage (unit, integration, E2E)
- **Accessibility:** Complete accessibility compliance and audits
- **Persistent Logs:** Add persistent notification logs/history
- **Mobile Support:** Add mobile and tablet support

---

## Glossary

- **SSE:** Server-Sent Events, a way for the server to push real-time updates to the client over HTTP
- **TanStack Query:** A library for fetching, caching, and updating server data in React apps
- **Container/Presentational Pattern:** A React pattern separating data-fetching/logic from UI rendering