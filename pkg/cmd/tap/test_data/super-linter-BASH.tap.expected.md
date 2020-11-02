BASH Linter
* [curlloop.sh](/curlloop.sh#L8) : 

```bash
    curl $APP 
         ^--^ SC2086  Double quote to prevent globbing and word splitting.

Did you mean  
    curl "$APP" 

For more information 
  https //www.shellcheck.net/wiki/SC2086 -- Double quote to prevent globbing ...
```
