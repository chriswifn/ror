{
  description = "Run or Raise for Hyprland";
  inputs.nixpkgs.url = "github.com/nixpkgs/nixos-unstable";
  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "ror";
        version = "0.0.1";
        src = ./.;
        vendorHash = null;
      };
      devShells.${system}.default = pkgs.mkShell { buildInputs = [ pkgs.go ]; };
    };
}
