package main

func DeepCopy() {
	// inString := "in"
	// outString := "out"
	inW := WorkloadIdentity{
		ServiceAccount: "in",
		Namespace:      "inNamespace",
	}

	println(&inW)
	println(&inW.ServiceAccount)

	var in *Wrapper = &Wrapper{
		Name: "inName",
		Work: []WorkloadIdentity{WorkloadIdentity{
			Namespace:      "test",
			ServiceAccount: "tset",
		}},
	}

	println(&(in.Work))

	println("---------------------")

	var out *Wrapper = new(Wrapper)
	println(&(out.Work))

	*out = *in

	println(out.Name)

	println(&out.Work)
	println(&in.Work)

	println("---------------------")

	println(&out)
	println(&in)

	println(&out.Name)
	println(&in.Name)

	println(out.Work[0].Namespace)

	println(&in.Work[0].ServiceAccount)
	println(&out.Work[0].ServiceAccount)

}

type Wrapper struct {
	Name string
	Work []WorkloadIdentity
}

type WorkloadIdentity struct {
	ServiceAccount string
	Namespace      string
}
