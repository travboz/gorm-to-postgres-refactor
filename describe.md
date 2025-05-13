## What do the directories in this project hold/represent?

Read this for more info:
https://tamerlan.dev/creating-go-rest-api/amp/#:~:text=username/go%2Dquest-,Folder%20Structure,-Model%2DView%2DController

Controllers – This component is responsible for handling requests from users.

Models – The central component of the pattern. It is the application's dynamic data structure, independent of the user interface. It directly manages the data, logic and rules of the application.

Utils – Contains helper functions that are used all over the project.

There is no `View` here because we're only building the backend API - no frontend 'view'.
