var data []string
// Declare second integer, double, and String variables.

scanner.Split(bufio.ScanLines)
for scanner.Scan() {
data = append(data, scanner.Text())
}
// Read and save an integer, double, and String to your variables.
i2, _ := strconv.ParseUint(data[0], 10, 64)
fmt.Println(i+i2)
// Print the sum of both integer variables on a new line.

d2, _ := strconv.ParseFloat(data[1], 64)
fmt.Printf("%.1f\n", (d+d2))
// Print the sum of the double variables on a new line.

fmt.Println(s+data[2])
// Concatenate and print the String variables on a new line
// The 's' variable above should be printed first.
