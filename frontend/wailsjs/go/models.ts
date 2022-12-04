export namespace ent {
	
	export class File {
	    id?: number;
	    // Go type: time.Time
	    create_time?: any;
	    // Go type: time.Time
	    update_time?: any;
	    src_path?: string;
	    current_path?: string;
	    filetype?: string;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.create_time = this.convertValues(source["create_time"], null);
	        this.update_time = this.convertValues(source["update_time"], null);
	        this.src_path = source["src_path"];
	        this.current_path = source["current_path"];
	        this.filetype = source["filetype"];
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
	export class Key {
	    id?: number;
	    // Go type: time.Time
	    create_time?: any;
	    // Go type: time.Time
	    update_time?: any;
	    type?: string;
	    public_key?: string;
	    private_key?: string;
	    references?: number;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.create_time = this.convertValues(source["create_time"], null);
	        this.update_time = this.convertValues(source["update_time"], null);
	        this.type = source["type"];
	        this.public_key = source["public_key"];
	        this.private_key = source["private_key"];
	        this.references = source["references"];
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
	    // Go type: time.Time
	    create_time?: any;
	    // Go type: time.Time
	    update_time?: any;
	    avatar?: string;
	    description?: string;
	    site_name?: string;
	    site_url?: string;
	    username?: string;
	    username_type?: string;
	    password?: string;
	    tags?: string[];
	
	    static createFrom(source: any = {}) {
	        return new PasswordItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.create_time = this.convertValues(source["create_time"], null);
	        this.update_time = this.convertValues(source["update_time"], null);
	        this.avatar = source["avatar"];
	        this.description = source["description"];
	        this.site_name = source["site_name"];
	        this.site_url = source["site_url"];
	        this.username = source["username"];
	        this.username_type = source["username_type"];
	        this.password = source["password"];
	        this.tags = source["tags"];
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

