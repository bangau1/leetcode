#[cfg(test)]
mod test {

    use std::cmp::{Ord, PartialOrd, Reverse};
    use std::collections::BinaryHeap;

    #[derive(Debug, PartialEq, Eq)]
    struct Vertex {
        node: u32,
        weight: i32,
    }

    impl PartialOrd for Vertex {
        fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
            return Some(self.cmp(other));
        }
    }

    impl Ord for Vertex {
        fn cmp(&self, other: &Self) -> std::cmp::Ordering {
            if self.weight == other.weight {
                return self.node.cmp(&other.node);
            }
            return self.weight.cmp(&other.weight);
        }
    }

    #[test]
    fn test_max_heap() {
        let mut max_heap = BinaryHeap::new();

        max_heap.push(1);
        max_heap.push(100);

        assert_eq!(100, max_heap.pop().unwrap());
    }

    #[test]
    fn test_min_heap_using_reverse() {
        let mut min_heap = BinaryHeap::new();

        min_heap.push(Reverse(1));
        min_heap.push(Reverse(100));

        assert_eq!(1, min_heap.pop().unwrap().0);
    }

    #[test]
    fn test_max_heap_on_struct() {
        let mut max_heap = BinaryHeap::new();

        max_heap.push(Vertex {
            node: 1,
            weight: 100,
        });
        max_heap.push(Vertex {
            node: 2,
            weight: 100,
        });
        max_heap.push(Vertex {
            node: 3,
            weight: 200,
        });

        assert_eq!(
            Vertex {
                node: 3,
                weight: 200
            },
            max_heap.pop().unwrap()
        );
        assert_eq!(
            Vertex {
                node: 2,
                weight: 100
            },
            max_heap.pop().unwrap()
        );
    }

    #[test]
    fn test_max_heap_on_struct_with_reverse() {
        let mut min_heap = BinaryHeap::new();

        min_heap.push(Reverse(Vertex {
            node: 1,
            weight: 100,
        }));
        min_heap.push(Reverse(Vertex {
            node: 2,
            weight: 100,
        }));
        min_heap.push(Reverse(Vertex {
            node: 3,
            weight: 200,
        }));

        assert_eq!(
            Vertex {
                node: 1,
                weight: 100
            },
            min_heap.pop().unwrap().0
        );
        assert_eq!(
            Vertex {
                node: 2,
                weight: 100
            },
            min_heap.pop().unwrap().0
        );
    }

    #[derive(Debug, PartialEq, Eq)]
    struct ReverseVertex(Vertex);

    impl PartialOrd for ReverseVertex {
        fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
            return Some(self.cmp(other));
        }
    }

    impl Ord for ReverseVertex {
        fn cmp(&self, other: &Self) -> std::cmp::Ordering {
            // we reverse the comparison from self->other to other->self
            return other.0.cmp(&self.0);
        }
    }

    #[test]
    fn test_max_heap_on_struct_with_wrapper_struct() {
        let mut min_heap = BinaryHeap::new();
        min_heap.push(ReverseVertex(Vertex {
            node: 1,
            weight: 100,
        }));
        min_heap.push(ReverseVertex(Vertex {
            node: 2,
            weight: 100,
        }));
        min_heap.push(ReverseVertex(Vertex {
            node: 3,
            weight: 200,
        }));

        assert_eq!(
            ReverseVertex(Vertex {
                node: 1,
                weight: 100
            }),
            min_heap.pop().unwrap()
        );
        assert_eq!(
            ReverseVertex(Vertex {
                node: 2,
                weight: 100
            }),
            min_heap.pop().unwrap()
        );
    }
}
