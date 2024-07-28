<?php
if (preg_match('#^/(theme|assets|js|css|images|resources)#i', $_SERVER["REQUEST_URI"]) === 1) {
  return false;
}

echo "Hello World! ";
