# auth-template

A reusable authentication and authorization template for system development.

This repository serves as a base template for implementing **multiple authentication and authorization strategies** using Go. Each strategy is developed in its own **phase and Git branch**, allowing clean separation, experimentation, and reuse across different projects.

---

## Frameworks Used / To Be Used

- **Go**

---

## Authentication & Authorization Phases

Each authentication or authorization method is implemented in a dedicated **phase/branch**. This makes it easy to explore, compare, or switch auth strategies without affecting others.

---

## Phase 1: JWT Authentication
**Branch:** `auth/jwt`

Stateless authentication using JSON Web Tokens (JWT). This phase covers token generation, validation, middleware protection, and secure request handling.

**Frameworks / Libraries:**
- JWT library (e.g. `golang-jwt`)
- Environment-based configuration

**Use Case:**
- REST APIs
- Microservices
- Stateless authentication systems

---

## Phase 2: RBAC (Role-Based Access Control)
**Branch:** `auth/rbac`

Authorization layer that controls access based on user roles and permissions. This phase focuses on protecting routes and actions after authentication.

**Frameworks / Libraries:**
- Custom RBAC middleware
- Database or in-memory role mapping

**Use Case:**
- Admin dashboards
- Enterprise systems
- Permission-based access control

---

## Phase 3: OAuth Authentication
**Branch:** `auth/oauth`

Third-party authentication using OAuth providers such as Google or GitHub. This phase handles external login flows and token exchange.

**Frameworks / Libraries:**
- OAuth2 libraries
- Third-party identity providers

**Use Case:**
- Social login
- External identity verification
- Consumer-facing applications

---

## Phase 4: Cookie-Based Authentication
**Branch:** `auth/cookie`

Session-based authentication using HTTP cookies. Focuses on session handling, secure cookie configuration, and server-side validation.

**Frameworks / Libraries:**
- Session management library
- Secure cookie handling

**Use Case:**
- Traditional web applications
- Server-rendered systems
- Long-lived user sessions

---

## Branching Strategy

Each authentication method is isolated in its own branch to enable:

- Clean separation of concerns
- Easy comparison between auth approaches
- Safe experimentation
- High reusability across projects

---

## Intended Use

- Backend authentication boilerplates
- Learning and experimentation
- Rapid prototyping
- Production-ready system foundations

---

## Status

This project is an evolving template. Additional authentication and authorization strategies may be added in future phases.
