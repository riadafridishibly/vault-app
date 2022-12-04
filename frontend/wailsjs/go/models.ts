export namespace ent {
	
	export class Key {
	    id?: number;
	    type?: string;
	    public_key?: string;
	    private_key?: string;
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type = source["type"];
	        this.public_key = source["public_key"];
	        this.private_key = source["private_key"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PasswordItem {
	    id?: number;
	    avatar?: string;
	    description?: string;
	    site_name?: string;
	    site_url?: string;
	    username?: string;
	    username_type?: string;
	    password?: string;
	    tags?: string[];
	    // Go type: time.Time
	    created_at?: any;
	    // Go type: time.Time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new PasswordItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.avatar = source["avatar"];
	        this.description = source["description"];
	        this.site_name = source["site_name"];
	        this.site_url = source["site_url"];
	        this.username = source["username"];
	        this.username_type = source["username_type"];
	        this.password = source["password"];
	        this.tags = source["tags"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

