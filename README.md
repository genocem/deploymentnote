# DeploymentNote

a simple cli tool to add notes to deployments in case you want to access some info fast

## Prerequisites

- **Go:** Version 1.16 or higher.
- **Kubernetes:** A working Kubernetes cluster with `kubectl` installed and configured.
- **Sudo Privileges:** Required for system-wide installation.

## Installation

Follow these steps to install DeploymentNote system-wide:

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/deploymentnote.git
   cd deploymentnote
   ```

2. **Build the Binary:**

   ```bash
   go build -o deploymentnote
   ```

3. **Make the Binary Executable:**

   ```bash
   chmod +x deploymentnote
   ```

4. **Move the Binary to `/usr/local/bin` (requires sudo):**

   ```bash
   sudo mv deploymentnote /usr/local/bin/
   ```

5. **Generate the Bash Completion Script:**

   ```bash
   go run ./ completion > deploymentnote_completion
   ```

6. **Install the Completion Script (requires sudo):**

   ```bash
   sudo mv deploymentnote_completion /etc/bash_completion.d/deploymentnote
   ```

7. **Reload Bash Completions:**

   ```bash
   source /etc/bash_completion.d/deploymentnote
   ```

## Usage

Once installed, simply run `deploymentnote` 

### Commands Overview
Both adding and deleting notes supports autocompletion of deployment name

- **Show Deployments:**

  Running `deploymentnote` without any subcommands displays a table of deployments with available custom notes:

  ```bash
  deploymentnote
  ```

- **Add a Note:**

  ```bash
  deploymentnote add deployment_name "write note here"
  ```

- **Delete a Note:**

  ```bash
  deploymentnote delete deployment_name
  ```


## Configuration

- **Custom Data File:**

  DeploymentNote uses a JSON file at `/tmp/custom_values.json` to store and retrieve custom notes. You may edit this file manually if needed.

- **Kubernetes Integration:**

  Deployment data is fetched via the command `kubectl get deployments`. Ensure your `kubectl` context is correctly set for your target cluster.
