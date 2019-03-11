# colorsort
Sort all 16 million RGB colors by luminance

Steps:
1) Create a slice of color structs with three nested for loops (0-255)
2) Run quicksort on the slice, but sort the colors by luminance
3) Write text file with all sorted colors

Luminance: `return c.r*0.5 + c.g*0.7 + c.b*0.2`
