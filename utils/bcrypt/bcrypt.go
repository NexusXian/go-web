package bcrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用bcrypt算法对密码进行哈希处理（使用默认工作因子）
// 输入原始密码字符串，返回哈希后的字符串和可能的错误
// HashPassword hashes a password using bcrypt with default cost
// Takes a raw password string and returns the hashed string and possible error
func HashPassword(password string) (string, error) {
	// 调用带工作因子的版本，使用bcrypt默认工作因子（通常为10）
	// Call the cost-specific version with bcrypt's default cost (usually 10)
	return HashPasswordWithCost(password, bcrypt.DefaultCost)
}

// HashPasswordWithCost 使用指定工作因子对密码进行哈希处理
// 工作因子范围：4-31（值越高，哈希强度越大但计算耗时越长）
// 输入原始密码和工作因子，返回哈希后的字符串和可能的错误
// HashPasswordWithCost hashes a password with specified cost factor
// Cost range: 4-31 (higher value means stronger hash but more computation time)
// Takes a raw password and cost factor, returns hashed string and possible error
func HashPasswordWithCost(password string, cost int) (string, error) {
	// 验证工作因子是否在有效范围内
	// Validate cost is within valid range
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return "", errors.New("invalid cost factor (must be between 4 and 31)")
	}

	// 使用指定工作因子生成哈希值
	// Generate hash with specified cost factor
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", errors.New("error hashing password")
	}

	return string(hashBytes), nil
}

// CheckPassword 验证原始密码与哈希值是否匹配
// 输入原始密码和哈希字符串，返回匹配结果（true表示匹配，false表示不匹配）
// CheckPassword verifies if a raw password matches a hashed password
// Takes a raw password and hashed string, returns match result (true if matched, false otherwise)
func CheckPassword(password, hash string) bool {
	// 比较原始密码与哈希值是否一致
	// Compare raw password with hashed value
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
