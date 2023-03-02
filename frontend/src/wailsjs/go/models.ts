export namespace Data {
	
	export class WorkerInfo {
	    requests_fulfilled: number;
	    kudos_rewards: number;
	    // Go type: struct { Generated float64 "json:\"generated\""; Uptime int "json:\"uptime\"" }
	    kudos_details: any;
	    performance: string;
	    threads: number;
	    uptime: number;
	    maintenance_mode: boolean;
	    nsfw: boolean;
	    trusted: boolean;
	    flagged: boolean;
	    uncompleted_jobs: number;
	    models: string[];
	    // Go type: struct { Name string "json:\"name\""; Id string "json:\"id\"" }
	    team: any;
	    bridge_agent: string;
	    max_length: number;
	    max_context_length: number;
	    type: string;
	    name: string;
	    id: string;
	    online: boolean;
	    owner?: string;
	    info?: string;
	
	    static createFrom(source: any = {}) {
	        return new WorkerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.requests_fulfilled = source["requests_fulfilled"];
	        this.kudos_rewards = source["kudos_rewards"];
	        this.kudos_details = this.convertValues(source["kudos_details"], Object);
	        this.performance = source["performance"];
	        this.threads = source["threads"];
	        this.uptime = source["uptime"];
	        this.maintenance_mode = source["maintenance_mode"];
	        this.nsfw = source["nsfw"];
	        this.trusted = source["trusted"];
	        this.flagged = source["flagged"];
	        this.uncompleted_jobs = source["uncompleted_jobs"];
	        this.models = source["models"];
	        this.team = this.convertValues(source["team"], Object);
	        this.bridge_agent = source["bridge_agent"];
	        this.max_length = source["max_length"];
	        this.max_context_length = source["max_context_length"];
	        this.type = source["type"];
	        this.name = source["name"];
	        this.id = source["id"];
	        this.online = source["online"];
	        this.owner = source["owner"];
	        this.info = source["info"];
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

