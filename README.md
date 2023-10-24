# datadiff

- package for showing differences in two json,yaml or any other data. Currently it supports two types of Comparers JsonComparer and YamlComparer
- Compare function:
  - Data of any type must implement an interface called Comparer that contains three definitions
    1.  IsEqual : You need to pass []byte data of two files. Lets say first file []byte is x and second file data is y. Then if both x and y are equal it returns true and nil (second return type is error.). If there is any error that is x is nil or y is nil then it returns false, error message. If x and y are not equal then it returns false, nil.

    2. AreEqual: Similar to IsEqual but rather than two input arguments, you can send variable number of arguments, all of them are []byte.Each []byte array is equal or not is compared only with the previous element.
    
    3. Compare: Takes two []byte arguments and return true if both of them are equal and if there are any differences then it returns newly added keys, deleted keys , keys are same but the data(values) are different keys , and error if there is any nil data or unmarshal errors.



  

  
