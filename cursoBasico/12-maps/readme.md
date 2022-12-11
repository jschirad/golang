This program demonstrates some common operations on maps in Go.

The main function starts by creating a new map called m using the make function. This map has string keys and int values. Then it adds two key-value pairs to the map, with the keys "Jose" and "Pepito" and the values 14 and 20 respectively.

Next, the program iterates over the elements in m using a for loop with range and prints each key-value pair. Then it retrieves the value associated with the key "Jose" and prints it.

Next, the program attempts to retrieve the value associated with the key "Juan", which does not exist in the map. Since the value is not found, the zero value for the value type (0 for int) is returned. This is printed to the screen.

Next, the program uses the ok idiom to check if a key exists in the map before trying to retrieve its value. If the key exists, ok will be true and the value will be printed. If the key does not exist, ok will be false and the zero value for the value type will be printed. This is demonstrated with the keys "Julio" and "Pepito".



