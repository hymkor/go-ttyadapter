- Added a `height` parameter to the callback function of the `Open` method.
  Originally, this package was developed as a subpackage of `nyaosorg/go-readline-ny`, where only the terminal width (`width`) was needed for single-line input.
  Since the package is now being redesigned as a standalone and more general-purpose library, it was necessary to include the terminal height (`height`) as well.

v0.0.1
------
Nov 1, 2025

- Initial release: separated functionality from github.com/nyaosorg/go-readline-ny
