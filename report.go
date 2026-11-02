package main

func (s *Store) totalpendapatan() float64{
	var Total float64
	for _,t := range s.Transaction{
		Total += t.Total
	}
	return Total
}


func (s *Store) barangterlaris() map[int]int{
	hasil := map[int]int{}
	for _,t := range s.Transaction{
		for _,item := range t.Items{
			hasil[item.IDbarang] += item.Jumlah
		}
	}
	return hasil
}
