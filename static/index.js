const buttons = [
  "(",
  ")",
  "%",
  "C",
  "7",
  "8",
  "9",
  "/",
  "4",
  "5",
  "6",
  "X",
  "1",
  "2",
  "3",
  "-",
  ".",
  "0",
  "=",
  "+",
];
const operators = [
  ["X", "/"],
  ["+", "-"],
];

input = document.querySelector("input");
answer = document.querySelector(".answer");
buttons.forEach((item) => {
  const div = document.createElement("div");
  div.classList.add("item");
  div.innerHTML = item;
  document.querySelector(".buttons").append(div);
});

document.querySelector(".buttons").addEventListener("click", (e) => {
  if (e.target.classList.contains("item")) {
    switch (e.target.innerHTML) {
      case "C":
        input.value = "";
        answer.innerHTML = "";
        break;
      case "=":
        console.log(calculate(input.value));
        break;
      case "X":
        input.value += "*"
        break;

      default:
        input.value += e.target.innerHTML;
        break;
    }
  }
});
const calculate = () => {
  answer.innerHTML = eval(input.value);
};

// const calculate = (string) => {
//   console.log("log", string);
//   let brackets = [];

//   if (string.indexOf("(") !== -1) {
//     for (let i = 0; i < string.length; i++) {
//       if (string[i] === "(") {
//         brackets.push(i);
//       } else if (string[i] === ")") {
//         const from = brackets.pop();
//         const res = calculate(string.slice(from + 1, i));
//       }
//     }
//   }

//   const priority = string.split("").map((el, i) =>
//     operators[0].indexOf(el) !== -1
//       ? "$" // первый приоритет
//       : operators[1].indexOf(el) !== -1
//       ? "#" // второй приоритет
//       : el
//   );
//   // .filter((a) => Array.isArray(a));
//   console.log(priority);

//   const divided = string.split(/([X+%/-])/);
//   console.log(divided);
//   console.log(divided.includes("X"));
//   while (divided.indexOf(/([X%/])/) !== -1) {
//     const index = divided.indexOf(/([X%/])/);
//     console.log(divided[index]);
//     switch (divided[index]) {
//       case "X":
//         console.log("hello");
//         const r = divided[index - 1] * divided[index + 1];
//         divided = divided.splice(divided[index - 1], 3, r);
//         break;
//       case "/":
//         const a = divided[index - 1] * divided[index + 1];
//         divided = divided.splice(divided[index - 1], 3, a);
//         break;
//       default:
//         break;
//     }
//   }
//   console.log(divided);
// };

// while (string.indexOf(/([X+%/-])/)) {
//   const index1 = string.indexOf(/([X%/])/);
//   const slice = string.slice(index1,  )
//   string = string.replace();
// }

// for (let i = 0; i < array.length; i++) {
//   const element = array[i];

// }

// const divided = oper.join("").split("!");
// const isSimpleProblem = divided.map((a) => Number.isInteger(+a));
// if (isSimpleProblem.indexOf(false) === -1) {
//   const arg1 = divided[0],
//     arg2 = divided[1];
//   switch (string[oper.indexOf("!")]) {
//     case "X":
//       return arg1 * arg2;
//       break;

//     default:
//       break;
//   }
// }
// oper
//
