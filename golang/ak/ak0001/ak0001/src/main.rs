use std::collections::HashMap;
pub struct Solution {}
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        //Vec

        let mut hsmap: HashMap<i32, usize> = HashMap::new();
        for (i, &num) in nums.iter().enumerate() {
            if let Some(&pos) = hsmap.get(&(target - num)) {
                return vec![pos as i32, i as i32];
            }
            hsmap.insert(num, i);
        }

        panic!()
    }
}
#[cfg(test)]
mod tests {
    use super::Solution;
    #[test]
    fn exploration() {
        println!("{:?}",Solution::two_sum(vec![2, 7, 11, 15], 9));
        assert_eq!(vec![1, 2], Solution::two_sum(vec![3, 2, 4], 6));

        //问题 :
        //1

    }
}

fn main() {}