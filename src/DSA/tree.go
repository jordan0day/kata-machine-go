package dsa

var Tree = BinaryNode[int]{
	Value: 20,
	Right: &BinaryNode[int]{
		Value: 50,
		Right: &BinaryNode[int]{
			Value: 100,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 30,
			Right: &BinaryNode[int]{
				Value: 45,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &BinaryNode[int]{
		Value: 10,
		Right: &BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

// export const tree: BinaryNode<number> = {
//     value: 20,
//     right: {
//         value: 50,
//         right: {
//             value: 100,
//             right: null,
//             left: null,
//         },
//         left: {
//             value: 30,
//             right: {
//                 value: 45,
//                 right: null,
//                 left: null,
//             },
//             left: {
//                 value: 29,
//                 right: null,
//                 left: null,
//             }
//         },
//     },
//     left: {
//         value: 10,
//         right: {
//             value: 15,
//             right: null,
//             left: null,
//         },
//         left: {
//             value: 5,
//             right: {
//                 value: 7,
//                 right: null,
//                 left: null,
//             },
//             left: null,
//         }
//     }
// };

var Tree2 = BinaryNode[int]{
	Value: 20,
	Right: &BinaryNode[int]{
		Value: 50,
		Right: nil,
		Left: &BinaryNode[int]{
			Value: 30,
			Right: &BinaryNode[int]{
				Value: 45,
				Right: &BinaryNode[int]{
					Value: 49,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
			Left: &BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left: &BinaryNode[int]{
					Value: 21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &BinaryNode[int]{
		Value: 10,
		Right: &BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

// export const tree2: BinaryNode<number> = {
//     value: 20,
//     right: {
//         value: 50,
//         right: null,
//         left: {
//             value: 30,
//             right: {
//                 value: 45,
//                 right: {
//                     value: 49,
//                     left: null,
//                     right: null,
//                 },
//                 left: null,
//             },
//             left: {
//                 value: 29,
//                 right: null,
//                 left: {
//                     value: 21,
//                     right: null,
//                     left: null,
//                 },
//             }
//         },
//     },
//     left: {
//         value: 10,
//         right: {
//             value: 15,
//             right: null,
//             left: null,
//         },
//         left: {
//             value: 5,
//             right: {
//                 value: 7,
//                 right: null,
//                 left: null,
//             },
//             left: null,
//         }
//     }
// };
