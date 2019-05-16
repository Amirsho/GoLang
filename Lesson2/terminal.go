package main
import "fmt"
import "time"
//проверка суммы
func sum_proverka(s string) int {
	var i, l, chek, chek2 int
	l = len(s)
	chek = 0
	chek2 = 0
	for i = 0; i < l; i++ {
		if s[i] == '.' {
			chek = chek + 1
		}
		if s[i] >= '0' && s[i] <= '9' {
			chek2 = chek2 + 1
		}
		if !(((s[i] >= '0') && (s[i] <= '9')) || (s[i] == '.')) {
			return 0
		}
	}
	if chek <= 1 && chek2 >= l-2 {
		return 1
	}
	return 0
}
//проверки номеров
func Megafon_proverka(a string, Megafon [10] string) int{
	var p string
	var i int
	for i = 0; i <= 8; i++ {
		if a[i] < '0' || a[i] > '9' {return 0}
	}
	p = a[0:3]
	for i = 0; i < 10; i++ {
		if p == Megafon[i] {return 1}
	}
	return 0
}
func Beeline_proverka(a string, Beeline [5] string) int{
	var p string
	var i int
	for i = 0; i <=8; i++{
		if a[i] < '0' || a[i] > '9'  {return 0}
	}
	p = a[0:3]
  	for i = 0; i < 4; i++ {
		if p == Beeline[i]  {return 1}
	  }
	  p = a[0:2]
  	  if p == Beeline[4]  {return 1}
  	  return 0
}
func Babilon_proverka(a string, Babilon [4] string) int{
	var p string
	var i int
	for i = 0; i <=8; i++{
		if a[i] < '0' || a[i] > '9' {return 0}
	}
	  p = a[0:3]  
	  for i = 0; i < 4; i++{
		if p == Babilon[i]  {return 1}
	  }
  	  return 0
}
func Tcell_proverka(a string, Tcell [3] string) int{
	var i int
	var p string
	for i = 0; i <=8; i++{
    if a[i] < '0' || a[i] > '9' {return 0}
	}
	p = "";
	p = a[0:3]
	for i = 0; i < 2; i++{
		if p == Tcell[i]  {return 1}
	}
		if p == Tcell[i] {return 1}
	p = a[0:2]
	if p == Tcell[2] {return 1}
	return 0
}

//проверка номеров!
func main(){
	var step = 1
	var operator_number, l int
	var t, Operator, number, Sum string
	Megafon := [10]string{	"900","901","902","903","904","905","906","909","908","880",}
	Beeline :=[5]string {"917","915","919","911","92",}
    Babilon :=[4]string {"918","988","985","987",};
	Tcell :=[3]string {"550","770", "93,"};
	fmt.Println(step)
	for step <  5 {

		if step == 1 {
			fmt.Println("Welcome to out pay terminal")
			fmt.Println("Choose oparation:")
			fmt.Println("1 - Choose operator")
			fmt.Println("esc - exit")
			fmt.Println("Back - Back")
			fmt.Scan(&t)
			if t == "1" {step = 2} else 
			if t == "esc" {step = 1
				continue} else
			if t == "Back" {
                step = step - 1
                if step == 0 {step = 1}
			} else {fmt.Println("Please, choose right oparation")}
			continue
		}

		
		if step == 2 {
			fmt.Println("Choose your operator:")
			fmt.Println("1 - Megafon")
			fmt.Println("2 - Beeline")
			fmt.Println("3 - Babilon")
			fmt.Println("4 - Tcell")
			fmt.Scan(&t)
			operator_number = int(t[0]) - 48;
			if t == "1"  {Operator = "Megafon"} else
			if t == "2"  {Operator = "Beeline"} else
			if t == "3"  {Operator = "Babilon"} else
			if t == "4"  {Operator = "Tcell"}   else
			if t == "esc" {step = 1
				continue} else
			if t == "Back" {
                step = step - 1
				if step == 0 {step = 1}
				continue
			} else{ 
				fmt.Println("Please, choose right operator number")
			 continue
			}
			 step = step + 1;
		 }

		 if step == 3 {
			fmt.Println("Enter your number(Operator):")
            fmt.Scan(&t)
            l = len(t);
            if t == "esc" {
                step = 1
                continue
           } else
            if t == "Back" {
                step = step - 1 
                if step == 0  {step = 1} 
            	continue
            }
            if l != 9  {
				fmt.Println("Please, enter right operator number")
                continue
			}
			
                if operator_number == 1{
                    if Megafon_proverka(t, Megafon) == 0{
					fmt.Println("Please, enter right operator number")
                    continue
                    }
				}else
                if operator_number == 2{
                    if Beeline_proverka(t, Beeline) == 0{
						fmt.Println("Please, enter right operator number")
						continue
                    }
                }else
                if operator_number == 3{
                    if Babilon_proverka(t, Babilon) == 0{
						fmt.Println("Please, enter right operator number")
                    continue
                    }
                }else
                if operator_number == 4{
                    if Tcell_proverka(t, Tcell) == 0 {
						fmt.Println("Please, enter right operator number")
                    continue
                    }
				}
            number = t
            step = step + 1
		}
		
		if step == 4 {
			fmt.Println("Enter sum:")
            fmt.Scan(&t)
			if t == "esc" {
                step = 1
                continue
           } else
            if t == "Back" {
                step = step - 1 
                if step == 0  {step = 1} 
            	continue
			}
            if sum_proverka(t) == 0 {
				fmt.Println("Please, enter right sum > 0 ")
                continue
			}
            Sum = t;
            if t[0] == '.'  {Sum = "0" + Sum}
            step = step + 1
        }
	}
	fmt.Println("------------------------------")
	fmt.Println("Operator:       ", Operator, operator_number)
	fmt.Println("Number:         ", number)
	fmt.Println("Sum:            ", Sum)
	fmt.Printf("Data:  ")
	dt := time.Now()
    fmt.Println(dt.Format("01-02-2006 15:04:05"))
	/*
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2019.01.02 21:13"))
	*/
	fmt.Println("OperationStatus: Succesful")
	fmt.Println("------------------------------")
}