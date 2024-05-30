# Jeeves: A Command Line LLM-based Question Answering Tool

## Overview

**Jeeves** is a command-line tool designed to provide answers to questions by leveraging different Large Language Model (LLM) APIs. Whether you need quick answers or detailed explanations, Jeeves is here to serve.

## Features (to be added)

- **Multiple LLM API Support**: Jeeves can interact with various LLM APIs to provide diverse and comprehensive answers.
- **Command Line Interface**: Simple and intuitive CLI for asking questions and receiving answers.
- **Customizable API Integration**: Easily configure and switch between different LLM APIs.
- **Extensible**: Designed to add new LLM APIs and functionality with ease.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- API keys for the LLM services you plan to use (e.g., OpenAI)

### Installation

1. **Install Jeeves**:

   ```sh
   go install github.com/igweo/jeeves@latest
   ```

2. **Set up your ENV with your api keys**:
   ```sh
   export OPENAI_API_KEY="my api key"
   ```

### Usage

To use Jeeves, simply run the executable with your question:

```sh
jeeves ask "When should I use a mutex over a channel?"
```
Or to use GPT4 run the command with the '--model' flag
```sh
jeeves ask --model="gpt4" "How do I do X in Y framework?"
```
