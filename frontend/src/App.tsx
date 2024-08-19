import { useState } from 'react';
import { DoCalculate } from '../wailsjs/go/main/App';
import './App.css';
import logo from './assets/images/logo-universal.png';

enum PROCESS_TYPE {
	ALGEBRAIC_BASIC,
	ALGEBRAIC_ADVANCED,
	RPN,
	MATHML,
}

const ProcessTypeLabel = [
	'Algebraic Basic',
	'Algebraic Advanced',
	'Reverse Polish Notation',
	'MathML',
];

const buttonNums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0, '.'];
function App() {
	const [resultText, setResultText] = useState('');
	const [processType, setProcessType] = useState(PROCESS_TYPE.ALGEBRAIC_BASIC);
	const [processString, setProcessString] = useState('');

	const captureInput = (input: string | number) => {
		setProcessString((old) => (old += String(input)));
	};

	const submitForCalculation = async () => {
		console.log('FROM UI', processString, processType);
		if (!processString) return;

		try {
			const res = await DoCalculate(processString, processType);
			setResultText(res);
		} catch (e) {
			console.error(e);
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

	return (
		<div id="App">
			<div>
				<div id="header">
					<select
						onChange={(e) =>
							setProcessType(e.target.value as unknown as PROCESS_TYPE)
						}
						defaultValue={PROCESS_TYPE.ALGEBRAIC_BASIC}
					>
						{Object.entries(PROCESS_TYPE).map(([i, val]) => {
							return (
								<option key={val} value={i}>
									{ProcessTypeLabel[+i]}
								</option>
							);
						})}
					</select>
				</div>
				<div id="result" className="result">
					{resultText}
				</div>
				<textarea
					id="raw-input"
					value={processString}
					onChange={(e) => setProcessString(e.target.value)}
				/>
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
