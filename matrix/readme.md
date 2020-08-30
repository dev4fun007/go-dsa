# Matrix

#### Transpose
* Square matrix transpose can be done in-place - a[i][j] swap with a[j][i] where j=i+1
* A rectangular matrix of MxN will transpose to another rectangular matrix of NxM - in-place not possible 

#### Rotation
O(n^2) time and O(1) space

* Clockwise 90 - Transpose and reverse each row
* Clockwise 180 - Reverse each row then reverse each column
* Anti-Clockwise 90 - Transpose and reverse each column
* Anti-Clockwise 180 - Reverse each column then reverse each row

