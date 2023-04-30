export namespace video {
	
	export class File {
	    data: number[];
	    ext: string;
	    id: string;
	    mimeType: string;
	    name: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.ext = source["ext"];
	        this.id = source["id"];
	        this.mimeType = source["mimeType"];
	        this.name = source["name"];
	        this.path = source["path"];
	    }
	}

}

