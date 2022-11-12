export namespace dao {
	
	export class PasswordItem {
	    description?: string;
	    site_name?: string;
	    username?: string;
	    password?: string;
	    tags?: string[];
	
	    static createFrom(source: any = {}) {
	        return new PasswordItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.site_name = source["site_name"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.tags = source["tags"];
	    }
	}

}

