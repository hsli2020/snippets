https://www.geeksforgeeks.org/getter-and-setter-in-python/

class Geeks:
    def __init__(self):
        self._age = 0
    
    # a getter function
    @property
    def age(self):
        print("getter method called")
        return self._age
    
    # a setter function
    @age.setter
    def age(self, a):
        if(a < 18):
            raise ValueError("Sorry you age is below eligibility criteria")
        print("setter method called")
        self._age = a

mark = Geeks()

mark.age = 19  # change property value triggers setter

print(mark.age)  # access property triggers getter

# mark.age = 17  # this triggers setter and raise error

""" output

setter method called
getter method called
19

"""
