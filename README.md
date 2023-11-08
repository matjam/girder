# girder

Girder is a package for Go that provides a general framework for implementing
roguelikes. Its mostly just for my own projects right now, but once the API
stabilizes I intend to stamp it as a 1.0 and provide support for it.

The package provides the following

 * A relatively simple ECS implementation for defining Entities, Components
   and Systems. If you're not interested in using an ECS, thats fine, but I
   wanted to provide one anyway as I'm using it for my own projects.

 * A generic 2d grid package that provides a way to store any type you want.
   This is useful when you have a "Tile" struct that you want to store in one
   grid, and then have another grid that stores a different type.

