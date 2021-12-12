//#include<math.h>

// pointer points to an integer at memory address 2048
int *data = (int*)2048;

float foo(float x)
{
    // set memory address
    *data = x + 1;

  //return cos(x);
  return 0.;
}