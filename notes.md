go build -o s2hWrapper.exe
LaucherBackground.exe s2hWrapper.exe -S "localhost:8080" -r "index.php"

ejemplo de index.php
```
<?php
if (preg_match('#^/(theme|assets|js|css|images|resources)#i', $_SERVER["REQUEST_URI"]) === 1) {
  return false;
}

echo "Hola Mundo";

```



