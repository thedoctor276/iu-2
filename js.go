package iu

import (
	"fmt"
	"strings"
)

const (
	WebkitBridge JSONBridge = "window.webkit.messageHandlers.onCallEventHandler.postMessage(JSON.stringify(msg));"
	EdgeBridge              = "alert('Edge as backend is not yet supported');"
	BlinkBridge             = "alert('Blink as backend is not yet supported');"

	frameworkJSTemplate = `
function MakeMouseEvent(event) {
    var obj = {
        "AltKey": event.altKey,
        "Button": event.button,
        "ClientX": event.clientX,
        "ClientY": event.clientY,
        "CtrlKey": event.ctrlKey,
        "Detail": event.detail,
        "MetaKey": event.metaKey,
        "PageX": event.pageX,
        "PageY": event.pageY,
        "ScreenX": event.screenX,
        "ScreenY": event.screenY,
        "ShiftKey": event.shiftKey,
    };

    return obj;
}

function MakeWheelEvent(event) {
    var obj = {
        "DeltaX": event.deltaX,
        "DeltaY": event.deltaY,
        "DeltaZ": event.deltaZ,
        "DeltaMode": event.deltaMode,
    };

    return obj;
}

function MakeKeyboardEvent(event) {
    var obj = {
        "AltKey": event.altKey,
        "CtrlKey": event.ctrlKey,
        "CharCode": event.charCode,
        "KeyCode": event.keyCode,
        "Location": event.location,
        "MetaKey": event.metaKey,
        "ShiftKey": event.shiftKey,
    };

    return obj;
}

function CallEventHandler(id, eventName, arg) {
    var argType = arg.type;
    
    switch (argType) {
        case "click":
        case "contextmenu":
        case "dblclick":
        case "mousedown":
        case "mouseenter":
        case "mouseleave":
        case "mousemove":
        case "mouseover":
        case "mouseout":
        case "mouseup":
        case "drag":
        case "dragend":
        case "dragenter":
        case "dragleave":
        case "dragover":
        case "dragstart":
        case "drop":
            arg = JSON.stringify(MakeMouseEvent(arg));
            break;
            
        case "wheel":
            arg = JSON.stringify(MakeWheelEvent(arg));
            break;
            
        case "keydown":
        case "keypress":
        case "keyup":
            arg = JSON.stringify(MakeKeyboardEvent(arg));
            break;

        default:
            arg = JSON.stringify(arg);
            break;
    };

    var msg = {
        "ID": id,
        "Name": eventName,
        "Arg": arg,
    };
    
    // JSON bridge
    %v
}
`
)

var (
	frameworkJS string
	jsonBridge  = WebkitBridge
)

type JSONBridge string

func init() {
	SetJSONBridge(WebkitBridge)
}

func SetJSONBridge(bridge JSONBridge) {
	var rawFrameworkJS string

	jsonBridge = bridge
	rawFrameworkJS = fmt.Sprintf(frameworkJSTemplate, jsonBridge)
	frameworkJS = strings.Trim(rawFrameworkJS, " \t\r\n")
}

func FrameworkJS() string {
	return frameworkJS
}
