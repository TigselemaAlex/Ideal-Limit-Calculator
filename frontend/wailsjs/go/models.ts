export namespace main {
	
	export class Frequency {
	    value: number;
	    count: number;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new Frequency(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.count = source["count"];
	        this.total = source["total"];
	    }
	}
	export class Word {
	    count: number;
	    frequency: number;
	
	    static createFrom(source: any = {}) {
	        return new Word(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count = source["count"];
	        this.frequency = source["frequency"];
	    }
	}
	export class Response {
	    words: {[key: string]: Word};
	    size: number;
	    frequencies: {[key: string]: Frequency};
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.words = this.convertValues(source["words"], Word, true);
	        this.size = source["size"];
	        this.frequencies = this.convertValues(source["frequencies"], Frequency, true);
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

