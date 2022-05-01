//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI
import Foundation

// MARK: - Scaff
struct Scaff: Codable {
    let scaffold: [Move]
    let toProjectID, fromProjectID: Int
}

// MARK: - Move
struct Move: Codable {
    let type: String
    let quantity: Int
}


struct TransfereScaffolding: View {
    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    @State private var confirmationMessage = ""
    @State private var showingConfirmation = false
    
    var body: some View {
        VStack {
            Text("Transfere scaffolding view")
            
            Button("Transfere scaffolding unit") {
                Task {
                    await transfereScaffoldingUnit()
                }
            }
            .padding()
        }
        .alert("Success!", isPresented: $showingConfirmation) {
            Button("OK") { }
        } message: {
            Text(confirmationMessage)
        }
    }

    /// https://www.appsdeveloperblog.com/http-post-request-example-in-swift/
    func transfereScaffoldingUnit() async {
        let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project/scaffolding")
        guard let requestUrl = url else { fatalError() }
        // Prepare URL Request Object
        var request = URLRequest(url: requestUrl)
        request.httpMethod = "PUT"
         
        
        let body : Scaff = Scaff(scaffold: [Move(type: "Bunnskrue", quantity: 3)], toProjectID: 4, fromProjectID: 12)
        
        guard let jsonData = try? JSONEncoder().encode(body) else {
            print("Failed to encode order")
            return
        }
        
        let jsonString = String(data: jsonData, encoding: .utf8)!
        
        request.httpBody = jsonString.data(using: String.Encoding.utf8);
        // Perform HTTP Request
        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
                
            // Check for Error
            if let error = error {
                print("Error took place \(error)")
                return
            }
     
            // Convert HTTP Response Data to a String
            if let data = data, let dataString = String(data: data, encoding: .utf8) {
                print("Response data string:\n \(dataString)")
            }
            
            if let response = response as? HTTPURLResponse {
                if response.statusCode == 200 {
                    // TODO: Need to remove or fix?
                    do {
                        let decodedOrder = try JSONDecoder().decode(Scaff.self, from: data!)
                        confirmationMessage = "You have successfully transfered \(decodedOrder.scaffold[0].quantity)x \(decodedOrder.scaffold[0].type) from project \(decodedOrder.fromProjectID) to project \(decodedOrder.toProjectID)!"
                    showingConfirmation = true
                    } catch {
                        print("\(response.statusCode) CODE")
                    }
                }
            }
        }
        task.resume()
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

struct ScaffoldingDetails: View {
    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool

    var body: some View {
        TransfereScaffoldingButton(scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        
        VStack {
            ScrollView(.vertical) {
                HStack {
                    Text("Left")
                    Divider()
                        .frame(width: 3, height: 400, alignment: .center)
                        .background(Color.gray)
                    Text("Right")
                }
            }
        }
    }
}

/*
struct TransfereScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffolding()
    }
}
*/



/*
 guard let url =  URL(string:"https://api.rangouts.com/signin")
         else{
             return
         }
         
         //### This is a little bit simplified. You may need to escape `username` and `password` when they can contain some special characters...
 let body : Welcome = Welcome(move: [Move(type: "Bunnskrue", quantity: 3)], toProject: 4, fromProject: 12)
 
 do {
     let jsonData = try JSONEncoder().encode(body)
     let jsonString = String(data: jsonData, encoding: .utf8)!
     print(jsonString)
     
     // and decode it back
     let decodedSentences = try JSONDecoder().decode([Welcome].self, from: jsonData)
     print(decodedSentences)
 } catch { print(error) }
 */
