# FableForge frontend

React (Create React App / `react-scripts`) UI for FableForge. See the
[repository README](../README.md) for the project overview and roadmap.

Current state: a Google sign-in landing page (`src/pages/Landing.jsx`) and a
placeholder home page (`src/pages/Home.jsx`) that previews the planned
chapter-to-illustrated-EPUB workflow. The UI does not call the backend yet.

## Setup

```sh
cp .env.example .env        # set REACT_APP_GOOGLE_CLIENT_ID
npm install
npm start                   # dev server on http://localhost:3000
npm run build               # production build into build/
```

## Environment

| Key | Purpose |
| --- | --- |
| `REACT_APP_GOOGLE_CLIENT_ID` | Google OAuth client ID for the sign-in button on the landing page |

Note: the landing page currently navigates to `/home` on any successful Google
sign-in; the credential is not yet verified by the backend.
