# go-cli-example

### build and install

- from the root directory  

```
go build  
go install  

# run  
go-cli-example add 5 2
```


### Possible path issues  

Mac M1 

- If after running go install the cli is not found, add the following to bash_profile or zshrc or wherever Path is updated:  

```
export PATH=$PATH:/Users/donquixote/go/bin
```
