# Use an official Node.js runtime as a parent image
FROM node:16

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Install additional necessary packages
RUN npm install @reduxjs/toolkit react-redux @mui/material @emotion/react @emotion/styled axios react-router-dom

# Copy the rest of the application code
COPY . .

# Build the React app for production
RUN npm run build

# Expose the port the app runs on
EXPOSE 3000

# Run the app
CMD ["npm", "start"]
