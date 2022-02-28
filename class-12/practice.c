#include <stdio.h>
int main()
{
    double a,b ,c, ave;

    scanf("%lf %lf %lf" , &a , &b , &c);
    ave = (a*2 + b*3 + c*5) / (2+3+5);
    printf("MEDIA = %.1lf\n" , ave);
    return 0;
}