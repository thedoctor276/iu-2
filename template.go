package iu

import "strings"

var (
	pageTemplate = strings.Trim(`
<!doctype html>
<html lang="{{.Lang}}">
<head>
    <title>{{.Title}}</title>
    <meta charset="utf-8" /> 
{{range .CSS}}
    <link rel="stylesheet" href="{{.}}" />{{end}}
</head>
<body>{{range .Body}}
{{.Render}}{{end}}
{{if .FrameworkJS}}
<script>
{{.FrameworkJS}}
</script>{{end}}
</body>
</html>
`, " \t\r\n")

	commonComponentTemplate = strings.Trim(`
<{{.Tag}} id="{{.ID}}"{{if .Class}}
     class="{{.Class}}"{{end}}{{if .Title}}
     title="{{.Title}}"{{end}}{{if .Lang}}
     lang="{{.Lang}}"{{end}}{{if .Spellcheck}}
     spellcheck="{{.Spellcheck}}"{{end}}{{if .Hidden}}
     hidden{{end}}{{if .Draggable}}
     draggable{{end}}{{if .ContentEditable}}
     contenteditable="{{.ContentEditable}}"{{end}}{{if .TabIndex}}
     tabindex="{{.TabIndex}}"{{end}}{{if .OnClick}}
     onclick="{{.OnEvent "OnClick" "event"}}"{{end}}{{if .OnContextMenu}}
     oncontextmenu="{{.OnEvent "OnContextMenu" "event"}}"{{end}}{{if .OnDblClick}}
     ondblclick="{{.OnEvent "OnDblClick" "event"}}"{{end}}{{if .OnMouseDown}}
     onmousedown="{{.OnEvent "OnMouseDown" "event"}}"{{end}}{{if .OnMouseEnter}}
     onmouseenter="{{.OnEvent "OnMouseEnter" "event"}}"{{end}}{{if .OnMouseLeave}}
     onmouseleave="{{.OnEvent "OnMouseLeave" "event"}}"{{end}}{{if .OnMouseMove}}
     onmousemove="{{.OnEvent "OnMouseMove" "event"}}"{{end}}{{if .OnMouseOver}}
     onmouseover="{{.OnEvent "OnMouseOver" "event"}}"{{end}}{{if .OnMouseOut}}
     onmouseout="{{.OnEvent "OnMouseOut" "event"}}"{{end}}{{if .OnMouseUp}}
     onmouseup="{{.OnEvent "OnMouseUp" "event"}}"{{end}}{{if .OnDrag}}
     ondrag="{{.OnEvent "OnDrag" "event"}}"{{end}}{{if .OnDragEnd}}
     ondragend="{{.OnEvent "OnDragEnd" "event"}}"{{end}}{{if .OnDragEnter}}
     ondragenter="{{.OnEvent "OnDragEnter" "event"}}"{{end}}{{if .OnDragLeave}}
     ondragleave="{{.OnEvent "OnDragLeave" "event"}}"{{end}}{{if .OnDragOver}}
     ondragover="{{.OnEvent "OnDragOver" "event"}}"{{end}}{{if .OnDragStart}}
     ondragstart="{{.OnEvent "OnDragStart" "event"}}"{{end}}{{if .OnDrop}}
     ondrop="{{.OnEvent "OnDrop" "event"}}"{{end}}{{if .OnScroll}}
     onscroll="{{.OnEvent "OnScroll" "event"}}"{{end}}{{if .OnWheel}}
     onwheel="{{.OnEvent "OnWheel" "event"}}"{{end}}{{if .OnKeyDown}}
     onkeydown="{{.OnEvent "OnKeyDown" "event"}}"{{end}}{{if .OnKeyPress}}
     onkeypress="{{.OnEvent "OnKeyPress" "event"}}"{{end}}{{if .OnKeyUp}}
     onkeyup="{{.OnEvent "OnKeyUp" "event"}}"{{end}}{{if .OnCopy}}
     oncopy="{{.OnEvent "OnCopy" "event"}}"{{end}}{{if .OnCut}}
     oncut="{{.OnEvent "OnCut" "event"}}"{{end}}{{if .OnPaste}}
     onpaste="{{.OnEvent "OnPaste" "event"}}"{{end}}{{if .OnBlur}}
     onblur="{{.OnEvent "OnBlur" "event"}}"{{end}}{{if .OnFocus}}
     onfocus="{{.OnEvent "OnFocus" "event"}}"{{end}}>{{range .Content}}
     {{.Render}}{{end}}</{{.Tag}}>
`, " \t\r\n")

	buttonTemplate = strings.Trim(`
<{{.Tag}} id="{{.ID}}"{{if .Class}}
     class="{{.Class}}"{{end}}{{if .Title}}
     title="{{.Title}}"{{end}}{{if .Lang}}
     lang="{{.Lang}}"{{end}}{{if .Spellcheck}}
     spellcheck="{{.Spellcheck}}"{{end}}{{if .Autofocus}}
     autofocus{{end}}{{if .Disabled}}
     disabled{{end}}{{if .Hidden}}
     hidden{{end}}{{if .Draggable}}
     draggable{{end}}{{if .ContentEditable}}
     contenteditable="{{.ContentEditable}}"{{end}}{{if .TabIndex}}
     tabindex="{{.TabIndex}}"{{end}}{{if .OnClick}}
     onclick="{{.OnEvent "OnClick" "event"}}"{{end}}{{if .OnContextMenu}}
     oncontextmenu="{{.OnEvent "OnContextMenu" "event"}}"{{end}}{{if .OnDblClick}}
     ondblclick="{{.OnEvent "OnDblClick" "event"}}"{{end}}{{if .OnMouseDown}}
     onmousedown="{{.OnEvent "OnMouseDown" "event"}}"{{end}}{{if .OnMouseEnter}}
     onmouseenter="{{.OnEvent "OnMouseEnter" "event"}}"{{end}}{{if .OnMouseLeave}}
     onmouseleave="{{.OnEvent "OnMouseLeave" "event"}}"{{end}}{{if .OnMouseMove}}
     onmousemove="{{.OnEvent "OnMouseMove" "event"}}"{{end}}{{if .OnMouseOver}}
     onmouseover="{{.OnEvent "OnMouseOver" "event"}}"{{end}}{{if .OnMouseOut}}
     onmouseout="{{.OnEvent "OnMouseOut" "event"}}"{{end}}{{if .OnMouseUp}}
     onmouseup="{{.OnEvent "OnMouseUp" "event"}}"{{end}}{{if .OnDrag}}
     ondrag="{{.OnEvent "OnDrag" "event"}}"{{end}}{{if .OnDragEnd}}
     ondragend="{{.OnEvent "OnDragEnd" "event"}}"{{end}}{{if .OnDragEnter}}
     ondragenter="{{.OnEvent "OnDragEnter" "event"}}"{{end}}{{if .OnDragLeave}}
     ondragleave="{{.OnEvent "OnDragLeave" "event"}}"{{end}}{{if .OnDragOver}}
     ondragover="{{.OnEvent "OnDragOver" "event"}}"{{end}}{{if .OnDragStart}}
     ondragstart="{{.OnEvent "OnDragStart" "event"}}"{{end}}{{if .OnDrop}}
     ondrop="{{.OnEvent "OnDrop" "event"}}"{{end}}{{if .OnScroll}}
     onscroll="{{.OnEvent "OnScroll" "event"}}"{{end}}{{if .OnWheel}}
     onwheel="{{.OnEvent "OnWheel" "event"}}"{{end}}{{if .OnKeyDown}}
     onkeydown="{{.OnEvent "OnKeyDown" "event"}}"{{end}}{{if .OnKeyPress}}
     onkeypress="{{.OnEvent "OnKeyPress" "event"}}"{{end}}{{if .OnKeyUp}}
     onkeyup="{{.OnEvent "OnKeyUp" "event"}}"{{end}}{{if .OnCopy}}
     oncopy="{{.OnEvent "OnCopy" "event"}}"{{end}}{{if .OnCut}}
     oncut="{{.OnEvent "OnCut" "event"}}"{{end}}{{if .OnPaste}}
     onpaste="{{.OnEvent "OnPaste" "event"}}"{{end}}{{if .OnBlur}}
     onblur="{{.OnEvent "OnBlur" "event"}}"{{end}}{{if .OnFocus}}
     onfocus="{{.OnEvent "OnFocus" "event"}}"{{end}}
     type="button">{{range .Content}}
     {{.Render}}{{end}}</{{.Tag}}>
`, " \t\r\n")

	inputTemplate = strings.Trim(`
<input id="{{.ID}}"{{if .Class}}
     class="{{.Class}}"{{end}}{{if .Title}}
     title="{{.Title}}"{{end}}{{if .Lang}}
     lang="{{.Lang}}"{{end}}{{if .Spellcheck}}
     spellcheck="{{.Spellcheck}}"{{end}}{{if .Hidden}}
     hidden{{end}}{{if .Autofocus}}
     autofocus{{end}}{{if .Checked}}
     checked{{end}}{{if .Disabled}}
     disabled{{end}}{{if .Draggable}}
     draggable{{end}}{{if .ContentEditable}}
     contenteditable="{{.ContentEditable}}"{{end}}{{if .TabIndex}}
     tabindex="{{.TabIndex}}"{{end}}{{if .Step}}
     step="{{.Step}}"{{end}}{{if .Min}}
     min="{{.Min}}"{{end}}{{if .Max}}
     max="{{.Max}}"{{end}}{{if .Value}}
     value="{{.Value}}"{{end}}
     type="{{.Type}}"{{if .OnClick}}
     onclick="{{.OnEvent "OnClick" "event"}}"{{end}}{{if .OnContextMenu}}
     oncontextmenu="{{.OnEvent "OnContextMenu" "event"}}"{{end}}{{if .OnDblClick}}
     ondblclick="{{.OnEvent "OnDblClick" "event"}}"{{end}}{{if .OnMouseDown}}
     onmousedown="{{.OnEvent "OnMouseDown" "event"}}"{{end}}{{if .OnMouseEnter}}
     onmouseenter="{{.OnEvent "OnMouseEnter" "event"}}"{{end}}{{if .OnMouseLeave}}
     onmouseleave="{{.OnEvent "OnMouseLeave" "event"}}"{{end}}{{if .OnMouseMove}}
     onmousemove="{{.OnEvent "OnMouseMove" "event"}}"{{end}}{{if .OnMouseOver}}
     onmouseover="{{.OnEvent "OnMouseOver" "event"}}"{{end}}{{if .OnMouseOut}}
     onmouseout="{{.OnEvent "OnMouseOut" "event"}}"{{end}}{{if .OnMouseUp}}
     onmouseup="{{.OnEvent "OnMouseUp" "event"}}"{{end}}{{if .OnDrag}}
     ondrag="{{.OnEvent "OnDrag" "event"}}"{{end}}{{if .OnDragEnd}}
     ondragend="{{.OnEvent "OnDragEnd" "event"}}"{{end}}{{if .OnDragEnter}}
     ondragenter="{{.OnEvent "OnDragEnter" "event"}}"{{end}}{{if .OnDragLeave}}
     ondragleave="{{.OnEvent "OnDragLeave" "event"}}"{{end}}{{if .OnDragOver}}
     ondragover="{{.OnEvent "OnDragOver" "event"}}"{{end}}{{if .OnDragStart}}
     ondragstart="{{.OnEvent "OnDragStart" "event"}}"{{end}}{{if .OnDrop}}
     ondrop="{{.OnEvent "OnDrop" "event"}}"{{end}}{{if .OnScroll}}
     onscroll="{{.OnEvent "OnScroll" "event"}}"{{end}}{{if .OnWheel}}
     onwheel="{{.OnEvent "OnWheel" "event"}}"{{end}}{{if .OnKeyDown}}
     onkeydown="{{.OnEvent "OnKeyDown" "event"}}"{{end}}{{if .OnKeyPress}}
     onkeypress="{{.OnEvent "OnKeyPress" "event"}}"{{end}}{{if .OnKeyUp}}
     onkeyup="{{.OnEvent "OnKeyUp" "event"}}"{{end}}{{if .OnCopy}}
     oncopy="{{.OnEvent "OnCopy" "event"}}"{{end}}{{if .OnCut}}
     oncut="{{.OnEvent "OnCut" "event"}}"{{end}}{{if .OnPaste}}
     onpaste="{{.OnEvent "OnPaste" "event"}}"{{end}}{{if .OnBlur}}
     onblur="{{.OnEvent "OnBlur" "event"}}"{{end}}{{if .OnFocus}}
     onfocus="{{.OnEvent "OnFocus" "event"}}"{{end}}{{if .OnSelect}}
     onselect="{{.OnEvent "OnSelect" "event"}}"{{end}}{{if .OnChanged }}
     onchange="{{.OnEvent "OnChanged" "this.value"}}"{{end}}{{if .OnChecked }}
     onchange="{{.OnEvent "OnChecked" "this.checked"}}"{{end}}>
`, " \t\r\n")
)

func PageTemplate() string {
	return pageTemplate
}

func CommonComponentTemplate() string {
	return commonComponentTemplate
}

func ButtonTemplate() string {
	return buttonTemplate
}

func InputTemplate() string {
	return inputTemplate
}
