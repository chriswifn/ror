# Run or Raise

Minimal implementation of run or raise functionality
for Hyprland using hyprctl dispatchers.

## Usage
```bash
ror <cmd> <class>
```

## Functionality
1. If the specified window `class` does not exist in the
   current window list, launch the application using `cmd`.
2. If there is one instance of the specified window `class`
   in the list of windows, focus that window.
3. If there are multiple instances of the specified window
   `class` cycle between them.
   
## Installation
1. Using Go: clone the repository and use `go install`
2. Using flake: Add the flake to the inputs section.
   Install the package as `inputs.ror.packages.${pkgs.system}.default`.
