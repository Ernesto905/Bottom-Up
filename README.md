# Bottoms Up 🍻 A Philosophical Framework Builder

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technical Stack](#technical-stack)
- [Site Structure](#site-structure)
- [License](#license)

## Overview

Bottoms Up is a single-page web application that allows users to construct a philosophical framework from the ground up. Users can choose from one of three branches of philosophy—Metaphysics, Epistemology, and Ethics—to begin their journey. The application guides users through a series of questions, offering real-time feedback and critiques through AI integration.

## Features

- **Topic Selection**: Choose from Metaphysics, Epistemology, or Ethics.
- **Tiered Questions**: Questions range from basic to advanced, adapting to user responses.
- **Question Types**: Mix of multiple-choice and open-ended questions.
- **AI Guidance**: Real-time feedback and critiques using OpenAI's API.
- **Results Analysis**: Detailed summary including potential biases, inconsistencies, and an overall score.

## Technical Stack

- **Backend**: Golang
- **Frontend**: HTMX
- **Database**: PostgreSQL
- **DevOps**: Docker, Kubernetes, Terraform
- **CI/CD**: GitHub Actions
- **AI**: OpenAI API

## Site Structure

### Home Page

- Introduction
- Hook and Description
- Disclaimer
- Call to Action
- Topic Selection Buttons

### Through the Wringer

- Introduction
- Topic Selection
- Question Types
- Questions

### Results

- Output Summary
- Potential Biases
- Glaring Issues
- Overall Score
- Things to Consider

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
