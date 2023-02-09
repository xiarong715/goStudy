package main

func main() {
	cc := CommonCounter{}
	mc := MutextCounter{}
	ac := AtomicCounter{}
	test(&cc)
	test(&mc)
	test(&ac)
}
