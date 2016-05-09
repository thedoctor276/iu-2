package kit

type Repeater struct {
	Source []DataContextComponent
}

func (r *Repeater) OnMount() {

}

func (r *Repeater) OnDismount() {

}

func (r *Repeater) Template() string {
	return `
<div class="Repeater">
    {{range .Source}}
        {{.Render}}
    {{end}}
</div>
`
}

func (r *Repeater) SetSource(src interface{}) {

}
