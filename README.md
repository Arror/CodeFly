# CodeFly

```thrift
namespace swift TT

enum Sex {
    male = 1
    female = 2
}

struct Person {
    1: required string name
    2: required i32 age
    3: required Sex sex
}

service Greet {

    string sayHallo(1: required Person person)
}
```	
```swift
// TTSex.swift
import Foundation

public enum TTSex: Int, Enum {
    
    public typealias E = Int
    
    case female = 2 
    case male = 1 
}

// TTPerson.swift
import Foundation

public struct TTPerson: Base {
    
    public var name: String?
    public var age: Int?
    public var sex: TTSex?

    public init?(json: Any?) {
        
        guard let dict = json as? [String: Any] else { return nil }
        
        name = dict <- "name"
        age = dict <- "age"
        sex = dict <- "sex"
    }
    
    public var json: Any {
        
        var dict = [String: Any]()
        
        dict["name"] = name?.json
        dict["age"] = age?.json
        dict["sex"] = sex?.json

        return dict
    }
}

// GreetService.swift
import Foundation

public struct GreetService {
    
    public static func sayHallo(person: TTPerson ,completion: @escaping (String) -> Void, failure: @escaping (Error) -> Void) -> Bool {

        guard let caller = Invokers.caller else { return false }

        let path = "Greet/sayHallo"
        
        var params = [String: Any]()
        
        params["person"] = person.json
        
        debugPrint("API: \(path)", "Request: ", params)
        
        caller.invoke(path: path, params: params, completion: { response in

            if let result = String(json: response) {

                debugPrint("API: \(path)", "Response: ", response)
                
                completion(result)

            } else {

                let error = InvokeError.invalidResponse
                
                debugPrint("API: \(path)", "Error: ", error)
                
                failure(error)
            }
            
        }, failure: { error in

            debugPrint("API: \(path)", "Error: ", error)
            
            failure(error)
        })

        return true
    }
    
}
```
	
	