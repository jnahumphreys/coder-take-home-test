# Product Requirements Document (PRD)

## Project Title
Coder Registry UI – Take-Home Assignment

## Purpose
To design and implement a customer-facing web UI for browsing, searching, and managing Coder Templates and Modules, as described in the provided backend API brief. This UI will demonstrate product thinking, user experience, and frontend engineering skills for the Coder hiring process.

## Background
Coder Templates define the infrastructure and configuration for cloud development environments. Modules are reusable components that extend templates with additional capabilities. The backend API provides endpoints to list, search, and manage these resources.

## Goals
- Provide a clear, intuitive interface for users to browse and search available Templates and Modules.
- Enable real-time awareness of new resources via notifications.
- Allow users to delete Modules (not Templates).
- Demonstrate best practices in UI/UX and frontend development.

## Out of Scope
- User authentication and authorization.
- Persistent notification logs (real-time only).
- Mobile and tablet support (desktop-only for this assignment).
- Full accessibility compliance (will document intended approach for a real-world scenario).

## User Stories & Requirements

### Common Requirements (All Pages)
- As a user, I want to see a list of resources (Templates or Modules) with key details:
  - Name
  - Description
  - Logo (image)
  - Contributor (GitHub username)
  - Operating System (Windows, Linux, MacOS)
  - Source (Partner, Official)
  - Custom tags (e.g., ai, gcp, aws)
- As a user, I want to search resources by name with auto-complete suggestions.
- As a user, I want to clear my search and filters easily.
- As a user, I want to receive real-time notifications (toast/banner) when a new resource is published, but only for the resource type I am currently viewing.

### Templates Page (`/templates`)
- As a user, I want to browse all available Templates.
- As a user, I want to search and filter Templates.
- As a user, I do not need to delete Templates.

### Modules Page (`/modules`)
- As a user, I want to browse all available Modules.
- As a user, I want to search and filter Modules.
- As a user, I want to delete a Module from the list.

## Non-Functional Requirements
- The UI should be visually appealing and consistent, using the Flowbyte design system (Tailwind CSS-based).
- The application should be easy to set up and run locally (devcontainer support).
- The UI should be responsive to desktop screen sizes.
- Accessibility: While not in scope for this assignment, the design should aim for clear contrast, keyboard navigation, and ARIA labeling in a real-world scenario.

## Success Criteria
- All requirements above are met and demonstrated in the UI.
- The application is easy to run and review.
- The UI is intuitive, visually consistent, and provides a positive user experience.
- Documentation is provided for setup, usage, and any trade-offs or future improvements.

## Future Considerations (Not in Scope)
- Full accessibility (WCAG) compliance.
- Mobile and tablet support.
- Persistent notification history.
- Integration with authentication and authorization systems.
