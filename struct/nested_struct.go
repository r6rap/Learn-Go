package structs

/*	Nested struct adalah anonymous struct yang diembed ke struct. 
	Deklarasinya langsung di dalam struct peng-embed */

type p1 struct{
	p2 struct{
		weapon []string
		level int
	}
	name string
}