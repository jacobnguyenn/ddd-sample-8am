# Entities layer

Entities contain business logic and interface through the other layers.
Business logic should be included in the model and service, and should not be depended any other layers. 
If we need to access any other layer, we should through the layers using repository interface. 
Inverting the dependencies like this, make packages to be isolated, more testable and maintainable.
