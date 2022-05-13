//
//  BlurView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 24/03/2022.
//

import SwiftUI

/// **BlurView**
/// A blur view
struct BlurView: UIViewRepresentable {
    let style: UIBlurEffect.Style
    
    func makeUIView(context: Context) -> UIVisualEffectView {
        let view = UIVisualEffectView(
            effect: UIBlurEffect(style: style)
        )
        
        return view
    }
    
    func updateUIView(_ uiView: UIVisualEffectView, context: Context) {
        // Do nothing here
    }
}

struct BlurView_Previews: PreviewProvider {
    static var previews: some View {
        BlurView(style: .systemMaterial)
    }
}
