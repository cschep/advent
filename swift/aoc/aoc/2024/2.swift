//
//  2.swift
//  aoc
//
//  Created by Christopher Schepman on 7/14/25.
//

import AppKit

func day2() {
    let app = NSApplication.shared
//    app.setActivationPolicy(.regular)
    
    let window = NSWindow(
        contentRect: NSMakeRect(200, 200, 800, 600),
        styleMask: [.titled, .closable, .resizable],
        backing: .buffered, defer: false
    )
    
    class DrawingView: NSView {
        override func draw(_ dirtyRect: NSRect) {
            guard let context = NSGraphicsContext.current?.cgContext else { return }
            
            // Your drawing code here
            context.setFillColor(NSColor.blue.cgColor)
            context.fill(CGRect(x: 100, y: 100, width: 200, height: 150))
        }
    }
    
    window.contentView = DrawingView(frame: window.contentView!.bounds)
    window.makeKeyAndOrderFront(nil)
    app.run()
}
