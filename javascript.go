package iu

import "fmt"

// Javascript bridges.
const (
	WebkitBridge JSBridge = "window.webkit.messageHandlers.onCallEventHandler.postMessage(JSON.stringify(msg));"
	EdgeBridge            = "alert('Edge as backend is not yet supported');"
	BlinkBridge           = "alert('Blink as backend is not yet supported');"
)

const (
	frameworkJSTemplate = `
function RenderComponent(id, component) {
    const sel = '[data-iu-id="' + id + '"]';
    const elem = document.querySelector(sel);
    elem.outerHTML = component;
}

function MakeMouseEvent(event) {
    const obj = {
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
    const obj = {
        "DeltaX": event.deltaX,
        "DeltaY": event.deltaY,
        "DeltaZ": event.deltaZ,
        "DeltaMode": event.deltaMode,
    };

    return obj;
}

function MakeKeyboardEvent(event) {
    const obj = {
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
    const argType = arg.type;
    
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
            
        case "mousewheel":
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
    
    const msg = {
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
	jsBridge    = WebkitBridge
)

// JSBridge represent a javascript snippet to communicate with a web browser.
// Should be only used in a driver implementation.
type JSBridge string

func init() {
	SetJSBridge(WebkitBridge)
}

// SetJSBridge set the bridge to use to communicate with a web browser.
// Should be only used in a driver implementation.
func SetJSBridge(bridge JSBridge) {
	jsBridge = bridge
	frameworkJS = fmt.Sprintf(frameworkJSTemplate, jsBridge)
}
