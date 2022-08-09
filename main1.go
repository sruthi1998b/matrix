package main
import ("fmt")
func main() {
    var rows int
    var columns int
    fmt.Print("Enter Rows Count: ")
    fmt.Scan(&rows)
	fmt.Print("Enter Columns Count: ")
    fmt.Scan(&columns)
    sample := [][]int{}
    
    //creating array
	for  i := 0; i < rows; i++ {
	  var row = make([]int, columns)
      sample = append(sample, row)
    }
    
    count := 1;
    for  k := 0; k < rows; k++ {
       sample[k][0]=count;
       count+=1;
        var i = k - 1;
        var j = 1;
        for !(i < 0 || i >= rows || j >= columns || j < 0) {
            sample[i][j]=count;
            count+=1;
            i--;
            j++;
        }
    }
    
    for  k := 1; k < columns; k++ {
        sample[rows - 1][k]=count;
        count+=1;
        i := rows - 2;
        j := k + 1;
        for !(i < 0 || i >= rows || j >= columns || j < 0) {
            sample[i][j]=count;
            count+=1;
            i--;
            j++;
        }
    }
   
    for _, row := range sample {
        
            fmt.Println(row)
    }
}
