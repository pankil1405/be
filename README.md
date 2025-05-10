# 🛠️ Be: A Collection of Golang Assertions

![GitHub release](https://img.shields.io/badge/releases-latest-blue.svg) [![GitHub issues](https://img.shields.io/github/issues/pankil1405/be.svg)](https://github.com/pankil1405/be/issues) [![GitHub forks](https://img.shields.io/github/forks/pankil1405/be.svg)](https://github.com/pankil1405/be/network/members) [![GitHub stars](https://img.shields.io/github/stars/pankil1405/be.svg)](https://github.com/pankil1405/be/stargazers)

Welcome to the **Be** repository! This project offers a wide collection of assertions for Golang, compatible with Gomega and Gomock. Whether you are writing tests for your Go applications or building testing libraries, Be provides a set of matchers to make your life easier.

## 📦 Getting Started

To get started with Be, you can download the latest release from the [Releases section](https://github.com/pankil1405/be/releases). Follow the instructions to download and execute the necessary files.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/pankil1405/be.git
   cd be
   ```

2. Install the dependencies:

   ```bash
   go get
   ```

3. Import Be in your Go files:

   ```go
   import "github.com/pankil1405/be"
   ```

### Usage

Be provides a variety of matchers for your testing needs. Here’s a quick example of how to use it:

```go
package main

import (
    "testing"
    . "github.com/pankil1405/be"
)

func TestExample(t *testing.T) {
    Expect(2 + 2).To(Equal(4))
}
```

This example shows a simple assertion that checks if 2 + 2 equals 4. You can use various matchers provided by Be to suit your testing requirements.

## 🛠️ Features

- **Gomega Compatibility**: Use Be seamlessly with Gomega matchers for Behavior-Driven Development (BDD).
- **Gomock Support**: Easily integrate with Gomock for mocking in tests.
- **Flexible Matchers**: A wide range of matchers to cover different testing scenarios.
- **Easy to Use**: Simple syntax that makes your tests clear and concise.

## 🔍 Matchers Overview

### Basic Matchers

- **Equal**: Checks if two values are equal.
- **NotEqual**: Checks if two values are not equal.
- **BeNil**: Checks if a value is nil.

### Collection Matchers

- **ContainElement**: Checks if a collection contains a specific element.
- **HaveLen**: Checks if a collection has a specific length.

### Error Matchers

- **HaveOccurred**: Checks if an error occurred.

### Custom Matchers

You can also create your own custom matchers to extend Be's functionality. Here’s how:

```go
func BeEven() types.GomegaMatcher {
    return &evenMatcher{}
}

type evenMatcher struct{}

func (m *evenMatcher) Match(actual interface{}) (success bool, err error) {
    num, ok := actual.(int)
    if !ok {
        return false, fmt.Errorf("expected an int, got %T", actual)
    }
    return num%2 == 0, nil
}

func (m *evenMatcher) FailureMessage(actual interface{}) string {
    return fmt.Sprintf("Expected %v to be even", actual)
}

func (m *evenMatcher) NegatedFailureMessage(actual interface{}) string {
    return fmt.Sprintf("Expected %v not to be even", actual)
}
```

## 📄 Documentation

For detailed documentation on all matchers and their usage, please refer to the [official documentation](https://github.com/pankil1405/be/docs).

## 🔗 Links

- [Releases](https://github.com/pankil1405/be/releases)
- [Issues](https://github.com/pankil1405/be/issues)
- [Contributing](https://github.com/pankil1405/be/blob/main/CONTRIBUTING.md)

## 🤝 Contributing

We welcome contributions! If you want to contribute to Be, please follow these steps:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

Please make sure to follow the coding standards and include tests for any new features.

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/pankil1405/be/blob/main/LICENSE) file for details.

## 🌟 Acknowledgments

We appreciate the contributions of the open-source community. Special thanks to the creators of Gomega and Gomock for their excellent libraries.

## 📅 Changelog

Please refer to the [CHANGELOG](https://github.com/pankil1405/be/blob/main/CHANGELOG.md) for a list of changes and updates.

## 📧 Contact

For questions or feedback, feel free to open an issue or contact us via the [issues page](https://github.com/pankil1405/be/issues).

---

Thank you for checking out the Be repository! We hope you find it useful for your Go testing needs. Don't forget to visit the [Releases section](https://github.com/pankil1405/be/releases) for the latest updates and features.