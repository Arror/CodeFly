# CodeFly
### Usage:

	Define .thrift file
	
	namespace swift TR

	enum TSex {
	    male = 1
	    female = 2
	}
	
	struct TUser {
	    1: required string name;
	    2: required i64 id;
	    3: required TSex sex
	}
	
	CodeFly will generate swift files
	
	// TRSex.swift
	import Foundation

	@objc
	public enum TRSex: Int, EnumItem {
	    
	    case male = 1 
	    case female = 2 
	}
	
	// TRUser.swift
	import Foundation

	public class TRUser: NSObject, JSONItem {
	    
	    var name: String?
	    var id: Int64?
	    var sex: TRSex?
	    
	    public var allKeys: Set<String> {
	        return ["name", "id", "sex"]
	    } 
	
	    public required init?(json: AnyObject?) {
	
	        super.init()
	
	        guard let json = json as? [String: AnyObject] else { return nil }
	        
	        self.name = String(json: dict["name"])
	        self.id = Int64(json: dict["id"])
	        self.sex = TRSex(json: dict["sex"])
	    }
	
	    public func toJSON() -> AnyObject? {
	
	        var json = [String: AnyObject]()
	        
	        json["name"] = self.name?.toJSON()
	        json["id"] = self.id?.toJSON()
	        json["sex"] = self.sex?.toJSON()
	
	        return json
	    }
	}
	
	