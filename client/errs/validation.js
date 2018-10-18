import ExtendableError from "./base";

class ValidationError extends ExtendableError {
	constructor(errors) {
		super(JSON.stringify(errors));
		this.errors = errors;
	}
}

export default ValidationError;
