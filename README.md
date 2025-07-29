# url-shortner
This repo contains the code related to url shortner
1. this repo has 3 layer of code 
	1. REST Interface which has exposed the functionlity of shortner , fetch original and metrics
		- here help of internet has been taken to know the converison lirbary and syntax 
	2. Shortner services which is internally handling of ShortenURL , OriginalURL and Metrics 
		- no internet help has been taken 
	3. Hashing mechanism 
		- written a small has function which every time multiply currenbt value with 31 and add new value of string characters
		- here no guarantee of clossion avoidance but still less collision  will occure
2. Docker file will be exposing the same Port for external interaction 
 
