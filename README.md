# Social 

## Folder structure
/bin - contains the binary executable (dist)
/cmd - entry points
/cmd/api - main server
/cmd/migrate - db migrations 
/internal - holds all internal packages (isolated from the entire app, not exported)
/docs - docs from swagger
/scripts - script for running server etc (misc/util/helper)
/web - contains the web (frontend)