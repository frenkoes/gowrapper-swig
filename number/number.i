 %module example
 %{
 /* Includes the header in the wrapper code */
 #include "number.h"
 %}
 
 /* Parse the header file to generate wrappers */
 %include "number.h"
 