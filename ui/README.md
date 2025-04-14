# SYNF - Front-End

## Overview

SYNF is a web application that provides real-time insights into device health, allowing users to monitor and manage system status from anywhere. This repository contains the front-end, built using React.

## Features

-   Responsive and modern UI
-   Real-time device health monitoring
-   Secure authentication and data handling
-   API integration with the SYNF backend

## Technologies Used

-   **React** - Component-based front-end framework
-   **React Router** - For client-side routing
-   **Vite** - For fast development and optimized builds

## Installation

### Prerequisites

Ensure you have the following installed on your system:

-   Node.js (>= 16.x)
-   npm or yarn

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/synf-ui.git
    cd synf-ui
    ```
2. Install dependencies:
    ```sh
    npm install
    # or
    yarn install
    ```
3. Start the development server:
    ```sh
    npm run dev
    # or
    yarn dev
    ```
4. Open your browser and go to `http://localhost:5173` (or as specified in the terminal output).

## Project Structure

```
├── public/              # Static assets
├── src/
│   ├── components/      # Reusable UI components
│   ├── pages/           # Application pages
│   ├── hooks/           # Custom React hooks
│   ├── context/         # Global state management
│   ├── services/        # API interaction and utilities
│   ├── App.tsx          # Main application entry point
│   ├── main.tsx         # Root render file
│   ├── index.css        # Global styles
├── package.json         # Dependencies and scripts
├── vite.config.ts       # Vite configuration
└── README.md            # Project documentation
```

## Deployment

1. Build the application:
    ```sh
    npm run build
    # or
    yarn build
    ```
2. Deploy the `dist/` folder to your preferred hosting service (e.g., Vercel, Netlify, or a self-hosted server).

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and submit a pull request.

## Contact

For inquiries or support, open an issue or reach out via nsrddyn@gmail.com
