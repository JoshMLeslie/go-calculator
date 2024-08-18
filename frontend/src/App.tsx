import { useState } from 'react';
import { Calculate } from '../wailsjs/go/main/App';
import './App.css';
import logo from './assets/images/logo-universal.png';

const useUpdateString: (
	init: string
) => [string, (v: string) => void, (v: string) => void] = (init: string) => {
	const [state, setState] = useState(init);
	const updateState = (newState: string) =>
		setState((oldState) => (oldState += newState));
	return [state, setState, updateState];
};

enum PROCESS_TYPE {
	ALGEBRAIC_BASIC = 'Algebraic Basic',
	ALGEBRAIC_ADVANCED = 'Algebraic Advanced',
	RPN = 'Reverse Polish Notation',
	MATHML = 'MathML',
}

const buttonNums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0, '.'];
function App() {
	const [processType, setProcessType] = useState(PROCESS_TYPE.ALGEBRAIC_BASIC);
	const [resultText, setResultText] = useState(0);
	const [processString, setProcessString, updateProcessString] =
		useUpdateString('');

	const captureInput = (input: string | number) => {
		if (input === '=') {
			console.log(input);
		} else {
			updateProcessString(String(input));
		}
	};

	const RenderOperators = () => {
		switch (processType) {
			case PROCESS_TYPE.ALGEBRAIC_BASIC:
				return (
					<div>
						<button onClick={() => captureInput('+')}>+</button>
						<button onClick={() => captureInput('-')}>-</button>
						<button onClick={() => captureInput('/')}>/</button>
						<button onClick={() => captureInput('*')}>*</button>
					</div>
				);
			case PROCESS_TYPE.RPN:
			case PROCESS_TYPE.ALGEBRAIC_ADVANCED:
				return (
					<div>
						<button onClick={() => captureInput('+')}>+</button>
						<button onClick={() => captureInput('-')}>-</button>
						<button onClick={() => captureInput('/')}>/</button>
						<button onClick={() => captureInput('*')}>*</button>
						<button onClick={() => captureInput('(')}>{'('}</button>
						<button onClick={() => captureInput(')')}>{')'}</button>
						<button onClick={() => captureInput('exp')}>exp</button>
						<button onClick={() => captureInput('sqr')}>sqr</button>
					</div>
				);
			case PROCESS_TYPE.MATHML:
				return <div>todo</div>;
		}
	};

	const submitForCalculation = () => {
		Calculate(processString).then(console.log);
	};

	return (
		<div id="App">
			<div>
				<div id="header">
					<select
						onChange={(e) => setProcessType(e.target.value as PROCESS_TYPE)}
						defaultValue={PROCESS_TYPE.ALGEBRAIC_BASIC}
					>
						{Object.values(PROCESS_TYPE).map((val) => {
							return (
								<option key={val} value={val}>
									{val}
								</option>
							);
						})}
					</select>
				</div>
				<div id="result" className="result">
					{resultText}
				</div>
				<div id="numeric-inputs">
					{buttonNums.map((num) => {
						return (
							<button
								key={'numeric-input-' + num}
								id={'button-num-' + num}
								onClick={() => captureInput(num)}
							>
								{num}
							</button>
						);
					})}
				</div>
				<div id="operator-inputs">
					<RenderOperators />
				</div>
				<button onClick={submitForCalculation}>Calculate</button>
			</div>
			<footer>
				<div>
					<p>Built with</p>
					<a
						href="https://github.com/wailsapp/wails"
						target="_blank"
                        rel="noopener noreferrer"
					>
						<img src={logo} id="logo" alt="wails logo" />
					</a>
				</div>
			</footer>
		</div>
	);
}

export default App;
